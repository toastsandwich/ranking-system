import { Link, Outlet } from "react-router-dom";
function App() {
  return (
    <div className="container">
      <div className="row">
        <div className="col">
          <h1 className="mt-3"> Welcome to realtime ranking service </h1>
          <hr className="mb-3" />
        </div>
      </div>
      <div className="row">
        <div className="col md-2">
          <nav>
            <div className="list-group">
              <Link to="/" className="list-group-item list-group-action">
                Submit
              </Link>
              <Link
                to="/get_rank"
                className="list-group-item list-group-action"
              >
                Get Rank
              </Link>
              <Link
                to="/list_top_n"
                className="list-group-item list-group-action"
              >
                List Top N
              </Link>
            </div>
          </nav>
        </div>
        <div className="col md-10">
          <Outlet />
        </div>
      </div>
    </div>
  );
}

export default App;
