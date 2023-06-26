import { createContext } from 'react';

interface FollowedSourcesContext {
  followedSources: ArticlesSourceInfoes;
  callAPIGetFollow: () => void;
}

export const FollowedSourcesContext = createContext<FollowedSourcesContext>({
  followedSources: [],
  callAPIGetFollow: () => {},
});
