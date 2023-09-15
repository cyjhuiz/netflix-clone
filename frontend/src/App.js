import "./App.css";
import MainNavBar from "./components/nav/MainNavBar";
import { Route, BrowserRouter as Router, Routes } from "react-router-dom";
import { AuthContext } from "./features/auth/context/auth-context";
import { Fragment, useContext } from "react";
import LandingPage from "./pages/LandingPage";
import HomePage from "./pages/HomePage";
import ShowPlayerPage from "./pages/ShowPlayerPage";
import MainLayout from "./layouts/MainLayout";

function App() {
  return (
    <div className="app">
      <Router>
        <Fragment>
          <Routes>
            <Route path="/" element={<LandingPage />} />
            <Route element={<MainLayout />}>
              <Route path="/browse" element={<HomePage />} />
            </Route>
            <Route
              path="/show/:showID/episode/:number"
              element={<ShowPlayerPage />}
            />
          </Routes>
        </Fragment>
      </Router>
    </div>
  );
}

export default App;
