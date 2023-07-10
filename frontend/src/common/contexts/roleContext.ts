import { createContext } from 'react';

interface RoleContext {
  role: UserRole;
  setRole: (userRole: UserRole) => void;
}

export const RoleContext = createContext<RoleContext>({
  role: {name: '', permissions: []},
  setRole: () => {},
});
