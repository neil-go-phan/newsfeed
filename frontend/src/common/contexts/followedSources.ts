import { createContext } from 'react';

interface FollowedSourcesContext {
  followedSources: ArticlesSourceInfoes;
  setFollowedSources: (articlesSources: ArticlesSourceInfoes) => void;
}

export const FollowedSourcesContext = createContext<FollowedSourcesContext>({
  followedSources: [],
  setFollowedSources: () => {},
});
