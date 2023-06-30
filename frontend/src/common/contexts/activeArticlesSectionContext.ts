import { createContext } from 'react';

export const SECTION_ALL_ARTICLES = 'all articles';
export const SECTION_UNREAD_ARTICLES = 'unread articles';
export const SECTION_READ_LATER_ARTICLES = 'read later articles';

interface ActiveSectionContextType {
  activeSection: string;
  setActiveSection: (newActiveSection: string) => void;
}

export const ActiveSectionContext = createContext<ActiveSectionContextType>({
  activeSection: SECTION_ALL_ARTICLES,
  setActiveSection: () => {},
});
