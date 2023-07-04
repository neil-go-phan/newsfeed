import React from 'react';
import LibReadlater from './libReadlater';
import LibRecentlyRead from './libRecentlyRead';

function Library() {
  return (
    <div className="library">
      <div className="library__title">Library</div>
      <LibReadlater />
      <LibRecentlyRead />
    </div>
  );
}

export default Library;
