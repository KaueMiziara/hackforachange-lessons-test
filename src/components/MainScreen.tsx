import React from "react";
import data from "../data/database.json";
import GradeSelector from "./GradeSelector";

function MainScreen() {
  return (
    <div>
        <h1>Main Menu</h1>
        <GradeSelector grades={data.grades} />
    </div>
  );
}

export default MainScreen;
