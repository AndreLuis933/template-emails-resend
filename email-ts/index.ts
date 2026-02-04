import { config } from "dotenv";

config();
const RESEND_API_KEY = process.env.RESEND_API_KEY;
const DESTINATION_EMAIL = process.env.DESTINATION_EMAIL;
const SENDER_EMAIL = process.env.SENDER_EMAIL;
const GITHUB_TEMPLATE_URL = process.env.GITHUB_TEMPLATE_URL;


export async function sendEmail(
  subject: string,
  htmlContent: string,
) {
  const response = await fetch("https://api.resend.com/emails", {
    method: "POST",
    headers: {
      Authorization: `Bearer ${RESEND_API_KEY}`,
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      from: SENDER_EMAIL,
      to: [DESTINATION_EMAIL],
      subject: subject,
      html: htmlContent,
    }),
  });

  return response.json();
}

async function loadTemplate(templateName: string): Promise<string> {
  const templateUrl = `${GITHUB_TEMPLATE_URL}/${templateName}`;
  const response = await fetch(templateUrl);

  if (!response.ok) {
    throw new Error(`Failed to fetch template: ${response.statusText}`);
  }

  return response.text();
}

export async function notifyError(jobName: string, error: Error): Promise<void> {
  const template = await loadTemplate("email_error.html");
  const timestamp = new Date()
    .toLocaleString("pt-BR", {
      day: "2-digit",
      month: "2-digit",
      year: "numeric",
      hour: "2-digit",
      minute: "2-digit",
      second: "2-digit",
      hour12: false,
    })
    .replace(",", " às");

  const errorMsg = error.stack || error.message;

  let html = template.replace("{{JOB_NAME}}", jobName);
  html = html.replace("{{TIMESTAMP}}", timestamp);
  html = html.replace("{{ERROR_MESSAGE}}", errorMsg);

  await sendEmail(`🚨 ${jobName} - Falhou`, html);
}

export async function notifySuccess(jobName: string): Promise<void> {
  const template = await loadTemplate("email_success.html");
  const timestamp = new Date()
    .toLocaleString("pt-BR", {
      day: "2-digit",
      month: "2-digit",
      year: "numeric",
      hour: "2-digit",
      minute: "2-digit",
      second: "2-digit",
      hour12: false,
    })
    .replace(",", " às");

  let html = template.replace("{{JOB_NAME}}", jobName);
  html = html.replace("{{TIMESTAMP}}", timestamp);

  await sendEmail(`✅ ${jobName} - Sucesso`, html);
}

async function main() {
  try {
    throw new Error("Erro proposital para testar notificação");
    await notifySuccess("Job Mensal Das");
  } catch (error) {
    await notifyError("Job Mensal Das", error as Error);
    throw error;
  }
}

if (import.meta.main) {
  main();
}
