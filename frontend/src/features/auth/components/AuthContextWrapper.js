import React from "react";
import { useAuth } from "../hooks/auth-hook";
import { AuthContext } from "../context/auth-context";

// provide auth related props to be accessed in App.js, specifically for React router with isLoggedIn prop
const AuthContextWrapper = ({ children }) => {
  const { userID, token, login, logout } = useAuth();

  return (
    <AuthContext.Provider
      value={{
        userID: userID,
        token: token,
        login: login,
        logout: logout,
        isLoggedIn: !!userID,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
};

export default AuthContextWrapper;
