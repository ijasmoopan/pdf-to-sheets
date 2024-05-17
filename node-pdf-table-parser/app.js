import { PdfDocument } from "@pomgui/pdf-tables-parser";
import fs from "fs";
import process from "process";
import * as path from "path";

console.log("Welcome to Pdf Parser by NODE..");
let pdfFileName = "company-details.pdf";
const jsonFileName = "tableData.json";
if (process.argv[2]) {
  pdfFileName = process.argv[2];
}
console.log("Reading pdf file:", pdfFileName);

const inputFilePath = path.join(process.cwd(), "../" + pdfFileName);
const outputFilePath = path.join(process.cwd(), "../" + jsonFileName);

console.log("Parsing pdf file...");

const pdf = new PdfDocument();
pdf
  .load(inputFilePath)
  .then(() =>
    fs.writeFileSync(outputFilePath, JSON.stringify(pdf, null, 2), "utf-8")
  )
  .catch((err) => console.error(err));

console.log("Parsing completed, created a new json file", jsonFileName);
console.log("Signing off from NODE!");
