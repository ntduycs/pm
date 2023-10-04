const removeDiacritics = (str?: string) => {
  if (!str) return '';

  return str.normalize('NFD').replace(/[\u0300-\u036f]/g, '');
};

export const Strings = {
  removeDiacritics,
};
