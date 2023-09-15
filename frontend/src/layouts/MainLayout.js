import React from "react";
import MainNavBar from "../components/nav/MainNavBar";
import { Outlet } from "react-router-dom";

const MainLayout = () => {
  return (
    <>
      <MainNavBar />
      <Outlet />
    </>
  );
};

export default MainLayout;
