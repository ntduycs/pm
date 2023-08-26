const fromObjects = (objects: object[]) => {
  return Object.values(objects).map((object) => Object.values(object));
};

export const Arrays = {
  fromObjects,
};
