import * as React from "react";
import Cookies from "js-cookie";
import axios from "axios";

export const useAuth = () => {
  const [token, setToken] = React.useState(undefined);
  const [userID, setUserID] = React.useState(undefined);

  const login = React.useCallback((userID, token) => {
    setUserID(userID);
    setToken(token);
    // update cookie store with login params
    let userData = {};
    const userDataCookie = Cookies.get("userData");
    if (userDataCookie !== undefined) {
      userData = JSON.parse(userDataCookie);
    }

    userData = {
      ...userData,
      userID,
      token,
    };

    Cookies.set("userData", JSON.stringify(userData));
  }, []);

  const logout = React.useCallback(() => {
    setUserID(undefined);
    setToken(undefined);

    Cookies.remove("userData");
  }, []);

  // initialize session on start up
  React.useEffect(() => {
    let userData = {};
    const userDataCookie = Cookies.get("userData");
    if (userDataCookie !== undefined) {
      userData = JSON.parse(userDataCookie);
      login(userData.userID, userData.token);
    }
  }, [login]);

  // add auth token to authorize requests
  axios.interceptors.request.use((config) => {
    // retrieve from cookie because can't access react state from authhook
    let token = "";
    const userDataCookie = Cookies.get("userData");
    if (userDataCookie !== undefined) {
      const userData = JSON.parse(userDataCookie);
      token = userData.token;
    }

    config.headers["Authorization"] = `Bearer ${token}`;
    return config;
  });

  return {
    userID,
    token,
    login,
    logout,
  };
};
