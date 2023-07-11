import React from 'react';
import { ThreeDots } from 'react-loader-spinner';

function OverlayLoading() {
  return (
    <div className="overlayLoading">
      <div className="text">
        <div className="threeDotLoading">
          <ThreeDots
            height="50"
            width="50"
            radius="9"
            color="#4fa94d"
            ariaLabel="three-dots-loading"
            visible={true}
          />
        </div>
        Loading...
      </div>
    </div>
  );
}

export default OverlayLoading;
