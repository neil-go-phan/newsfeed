import { _ROUTES } from '@/helpers/constants';
import Link from 'next/link';
import React from 'react';

function Banner() {
  return (
    <div className="banner">
      <div className="banner__content">
        <h1>Build your own newsfeed</h1>
        <p>Ready to give it a go?</p>
        <Link
          className="btn btn-primary d-block px-3 text-nowrap"
          href={_ROUTES.REGISTER_PAGE}
        >
          <span>Create account</span>
        </Link>
      </div>
    </div>
  );
}

export default Banner;
