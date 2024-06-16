import { Route, RouterProvider, createBrowserRouter, createRoutesFromElements } from "react-router-dom"
import Dashboard from "./Pages/Dashboard";
import Races from "./Pages/Races";
import Shell from "./Components/Shell";
import LeaderBoard from "./Pages/LeaderBoard";

const router = createBrowserRouter(
  createRoutesFromElements(
    <Route path="/" element={<Shell />}>
      <Route path="/" element={<Dashboard />} />
      <Route path="races" element={<Races />} />
      <Route path="leader-board" element={<LeaderBoard />} />
      {/* ... etc. */}
    </Route>
  )
);


function App() {

  return (
    <RouterProvider router={router} />
  )
}

export default App
