import { createContext } from 'react';

interface CategoriesContext {
  categories: Categories;
  setCategories: (categories: Categories) => void;
}

export const CategoriesContext = createContext<CategoriesContext>({
  categories: [],
  setCategories: () => {},
});
