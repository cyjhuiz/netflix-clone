import React, { useContext, useState } from "react";
import { useNavigate } from "react-router-dom";
import { AuthContext } from "../features/auth/context/auth-context";
import LandingPageNavBar from "../components/nav/LandingPageNavBar";
import SignUpLandingForm from "../features/auth/components/SignUpLandingForm";
import SignUpModal from "../features/auth/components/SignUpModal";
import LoginModal from "../features/auth/components/LoginModal";

const LandingPage = () => {
  const { isLoggedIn } = useContext(AuthContext);
  const [mode, setMode] = useState("LANDING_PAGE");
  const [email, setEmail] = useState("");

  const navigate = useNavigate();
  if (isLoggedIn) {
    navigate("/browse");
  }

  return (
    <div
      style={{
        backgroundImage: `url("/images/landing-page-background.jpg")`,
        height: "100vh",
        minHeight: "100vh",
        backgroundSize: "cover",
        zIndex: "-2",
      }}
    >
      <div
        id="backdrop"
        style={{
          width: "100%",
          zIndex: "-1",
          height: "100vh",
          background: "rgba(0, 0, 0, 0.4)",
        }}
      >
        <LandingPageNavBar setMode={setMode} />
        {mode === "LANDING_PAGE" ? (
          <SignUpLandingForm
            email={email}
            setEmail={setEmail}
            setMode={setMode}
          />
        ) : mode === "SIGN_UP" ? (
          <SignUpModal initialInputEmail={email} setMode={setMode} />
        ) : (
          <LoginModal setMode={setMode} />
        )}
      </div>
    </div>
  );
};

export default LandingPage;
