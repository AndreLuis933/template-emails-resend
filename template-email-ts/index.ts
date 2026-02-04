import { readFileSync } from "fs";
import { join, dirname } from "path";
import { config } from "dotenv";

config();
const RESEND_API_KEY = process.env.RESEND_API_KEY;
const DESTINATION_EMAIL = process.env.DESTINATION_EMAIL;
const SENDER_EMAIL = process.env.SENDER_EMAIL;


async function sendEmail(
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

function loadTemplate(templateName: string): string {
  const templatePath = join(dirname(import.meta.path), "..", "template", templateName);
  return readFileSync(templatePath, "utf-8");
}

async function notifyError(jobName: string, error: Error): Promise<void> {
  const template = loadTemplate("email_error.html");
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

async function notifySuccess(jobName: string): Promise<void> {
  const template = loadTemplate("email_success.html");
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
    const a = 1 / 0;
    await notifySuccess("Job Mensal Das");
  } catch (error) {
    await notifyError("Job Mensal Das", error as Error);
    throw error;
  }
}

if (import.meta.main) {
  main();
}
