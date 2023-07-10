import { _ROUTES } from '@/helpers/constants';
import Link from 'next/link';
import React from 'react';

function IntroducionParagraph() {
  return (
    <div className="landing__introductionParagraph">
      <div className="mt-0 flex flex-column d-sm-block hero-bgr col-12 mx-auto">
        <h1 className="mt-0 mt-sm-4 mb-4 mb-sm-0">Build your own newsfeed</h1>
        <p className="py-3 px-lg-5 px-0 py-sm-4 mt-2 mx-auto">
          Stay informed with tailored news based on your interests and
          preferences. Our platform offers diverse sources, customization
          options, and a user-friendly interface. Trust in our
          commitment to accuracy, quality, and ethical journalism. Join us today
          for a personalized news experience.
        </p>
        <div className="landing__introductionParagraph--buttons pb-2 mb-2 pb-sm-5 mb-sm-5">
          <Link
            className="registerBtn btn btn-primary mb-3"
            href={_ROUTES.REGISTER_PAGE}
          >
            Create account
          </Link>
          <Link
            className="featuresBtn btn-tertiary btn ms-sm-2 ms-lg-3 mb-3"
            href={_ROUTES.DISCOVER_PAGE}
          >
            Discover our servies
          </Link>
        </div>
        <div className="img position-relative mb-5 w-100">
          <div className="h-100">
            <img alt="example" src={'/images/examplefeed.png'} />
          </div>
        </div>
      </div>
    </div>
  );
}

export default IntroducionParagraph;
