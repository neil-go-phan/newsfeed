import { faCompass } from '@fortawesome/free-regular-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import React from 'react';
import CompareItem from './compareItem';

export type CompareFeatureItem = {
  name: string;
  description: string;
  basic: boolean;
  premium: boolean;
};

export const compareItems: Array<CompareFeatureItem> = [
  {
    name: 'Follow feeds',
    description: 'Follow websites feeds.',
    basic: true,
    premium: true,
  },
  {
    name: 'Trending articles',
    description: 'See what is currently trending among our users.',
    basic: true,
    premium: true,
  },
  {
    name: 'Search in feeds',
    description: 'Search for articles previously collected by your feeds.',
    basic: true,
    premium: true,
  },
  {
    name: 'Read later list',
    description:
      'Save articles to read later and collect all starred items in your library.',
    basic: true,
    premium: true,
  },
  {
    name: 'Add new source',
    description: 'Can not find your favorited feeds ? Add new now!',
    basic: false,
    premium: true,
  },
];

function CompareFeature() {
  return (
    <>
      <div className="container-xxl py-5 px-3 px-sm-5">
        <h2 className="pt-5 heading-h2">Compare features</h2>
      </div>
      <div className="text-left container-fluid">
        <div className="row py-4 pricing__feature">
          <div className="container-xxl px-2 px-md-5 py-3">
            <div className="row mx-1 px-0">
              <div className="col-4 d-flex align-items-start d-none d-md-flex">
                <FontAwesomeIcon
                  className="d-none d-md-block icon"
                  icon={faCompass}
                />
                <div className="mx-md-3">
                  <h5 className="heading-h5 mb-0">Features</h5>
                  <p className="prg-sm d-none d-md-block">
                    Get the best from the web in a single place
                  </p>
                </div>
              </div>
              <div className="col-sm-8">
                <div className="d-flex">
                  <div className="col heading-h5 text-center">Basic</div>
                  <div className="col heading-h5 text-center">Premium</div>
                </div>
              </div>
            </div>
          </div>
          <div className="pricing__feature--rows">
            <div className="container-xxl px-3 px-sm-5">
              {compareItems.map((item) => (
                <CompareItem item={item} />
              ))}
            </div>
          </div>
        </div>
      </div>
    </>
  );
}

export default CompareFeature;
