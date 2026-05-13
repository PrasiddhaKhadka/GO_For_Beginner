type Translation = {
  en: string;
  ja: string;
};

export const Translations: Record<string, Translation> = {
  ADMIN_CANNOT_APPLY_FOR_CREATOR: {
    en: "Admin cannot apply for creator access",
    ja: "管理者はクリエイターとして申請できません",
  },
  HELLO: {
    en: "Hello, welcome!",
    ja: "こんにちは、ようこそ！",
  },
  GREETING: {
    en: "Good morning",
    ja: "おはようございます",
  },
};
