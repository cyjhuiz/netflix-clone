import { createContext } from "react";

export const AuthContext = createContext({
  userID: undefined,
  token: undefined,
  login: () => {},
  logout: () => {},
});
