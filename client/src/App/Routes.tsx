import React from "react";
import {
  Switch,
  Route
} from "react-router-dom";
import Patients from "./Patients/Patients"

// This site has 3 pages, all of which are rendered
// dynamically in the browser (not server rendered).
//
// Although the page does not ever refresh, notice how
// React Router keeps the URL up to date as you navigate
// through the site. This preserves the browser history,
// making sure things like the back button and bookmarks
// work properly.

export default function Routes() {
  return (
    <Switch>
        <Route exact path="/">
            <Patients />
        </Route>
        {/* <Route path="/dashboard">
            <Dashboard />
        </Route> */}
    </Switch>
  );
}