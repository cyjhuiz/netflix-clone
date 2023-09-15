import React, { useContext } from "react";
import ShowCarousell from "../features/show-carousell/components/ShowCarousell";
import FeaturedShowBanner from "../features/featured-show-banner/components/FeaturedShowBanner";
import { AuthContext } from "../features/auth/context/auth-context";
import { useNavigate } from "react-router-dom";

const HomePage = () => {
  const categories = [
    "Netflix Originals",
    "Trending Now",
    "Top Rated",
    "Action Movies",
    "Comedy Movies",
    "Horror Movies",
    "Romance Movies",
    "Documentaries",
  ];

  const { isLoggedIn } = useContext(AuthContext);
  const navigate = useNavigate();
  if (!isLoggedIn) {
    navigate("/");
  }

  return (
    <React.Fragment>
      <FeaturedShowBanner />

      <div>
        {categories.map((category, index) => (
          <ShowCarousell key={`showCarousell_${index}`} category={category} />
        ))}
      </div>
    </React.Fragment>
  );
};

export default HomePage;
