import React from 'react';

function RemarkableFeatures() {
  return (
    <div className="landing__remarkableFeatures py-lg-5">
      <div className="info-section">
        <div className="d-flex justify-content-between flex-column flex-lg-row">
          <div className="img order-2 order-md-1 feature-row-image position-relative flex flex-column flex-md-row justify-content-center col-12 col-md-8 col-lg-6"></div>
          <div className="content col-12 col-md-8 col-lg-6 text-left order-1 order-md-2">
            <h3 className="heading-h3">
              Follow your favorite websites and creators
            </h3>
            <div className="prg-md text-left mt-3 mt-md-2">
              Bring the content that matters to you together and enjoy the best
              from the web in a single place.
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default RemarkableFeatures;
