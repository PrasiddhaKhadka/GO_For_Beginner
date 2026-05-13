import fs from "fs";
import { Translations } from "./translations.ts";

type TranslationObject = {
  other: string;
};

type TranslationFile = Record<string, TranslationObject>;

const getTranslationFiles = () => {
  let enTranslationFile: TranslationFile = {};
  let jaTranslationFile: TranslationFile = {};

  for (const key in Translations) {
    const t = Translations[key];
    if (t) {
      enTranslationFile[key] = { other: t.en };
      jaTranslationFile[key] = { other: t.ja };
    }
  }
  return { enTranslationFile, jaTranslationFile };
};

async function generateCode() {
  const keys = Object.keys(Translations);

  // Generate constants.go
  const code = `package i18n
    const(
    ${keys.map((k) => `\t${k} TranslationKey = "${k}"`).join("\n")}
    )
  `;
  fs.writeFileSync("../constants.go", code);

  // Generate en.json and ja.json
  const { enTranslationFile, jaTranslationFile } = getTranslationFiles();
  fs.writeFileSync("../en.json", JSON.stringify(enTranslationFile, null, 2));
  fs.writeFileSync("../ja.json", JSON.stringify(jaTranslationFile, null, 2));

  console.log("Generated: constants.go, en.json, ja.json");
}

generateCode();
