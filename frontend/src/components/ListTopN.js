const ListTopN = () => {
  return (
    <>
      <div className="container mt-2">
        <h2 className="text-center">Leader Board</h2>
        <hr />
        <form className="row gx-2 gy-3">
          {/* Location and Number of Top Ranks inputs on the same line */}
          <div className="col-md-6">
            <label htmlFor="location" className="form-label">
              Location
            </label>
            <input
              type="text"
              className="form-control"
              id="location"
              placeholder="Enter location"
            />
            <div className="form-text">
              If you want global ranking, enter "global".
            </div>
          </div>
          <div className="col-md-6">
            <label htmlFor="n" className="form-label">
              Number of Top Ranks
            </label>
            <input type="number" className="form-control" id="n" />
          </div>

          {/* Radio buttons for selecting ranking type and Submit button below */}
          <div className="col-12 d-flex align-items-center">
            <label className="form-label me-3 mb-0">Rank By:</label>
            <div className="form-check me-3">
              <input
                className="form-check-input"
                type="radio"
                name="rankBy"
                id="rankByCountry"
                value="country"
                defaultChecked
              />
              <label className="form-check-label" htmlFor="rankByCountry">
                Country
              </label>
            </div>
            <div className="form-check me-3">
              <input
                className="form-check-input"
                type="radio"
                name="rankBy"
                id="rankByState"
                value="state"
              />
              <label className="form-check-label" htmlFor="rankByState">
                State
              </label>
            </div>
            <button type="submit" className="btn btn-primary btn-sm ms-auto">
              Submit
            </button>
          </div>
        </form>
      </div>
    </>
  );
};

export default ListTopN;
