import { Link, Outlet } from "react-router-dom";

export default function Shell() {
  return (
    <div>
        <Link to="/">Home</Link>
        <Link to="/races">Races</Link>
        <Link to="/leader-board">Leader Board</Link>
        <Outlet/>
    </div>
  );
}