import { createContext } from 'react';

interface SearchKeywordContext {
  keyword: string;
  setKeyword: (keyword: string) => void;
}

export const SearchKeywordContext = createContext<SearchKeywordContext>({
  keyword: '',
  setKeyword: () => {},
});
