import { createContext } from 'react';


interface TriggerRefreshContextType {
  triggerRefresh: boolean;
  setTriggerRefresh: (triggerRefresh: boolean) => void;
}

export const TriggerRefreshContext = createContext<TriggerRefreshContextType>({
  triggerRefresh: false,
  setTriggerRefresh: () => {},
});
