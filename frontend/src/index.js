import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";
import SubmitForm from "./components/SubmitForm";
import GetRank from "./components/GetRank";
import ListTopN from "./components/ListTopN";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import ErrorPage from "./components/ErrorPage";

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    errorElement: <ErrorPage />,
    children: [
      {
        index: true,
        element: <SubmitForm />,
      },
      {
        path: "/get_rank/",
        element: <GetRank />,
      },
      {
        path: "/list_top_n/",
        element: <ListTopN />,
      },
    ],
  },
]);
const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>,
);
