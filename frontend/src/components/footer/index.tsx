import { Feed } from '@mui/icons-material';
import { Button, Typography } from '@mui/material';
import React from 'react';
import { pages } from '../navbar';
import Link from 'next/link';

const GITHUB_LINK = 'https://github.com/neil-go-phan/newsfeed';

function FooterComponent() {
  return (
    <div className="footerComponent py-5">
      <div className="footerComponent__logo">
        <Feed sx={{ mr: 1, color: 'white' }} />
        <Typography
          variant="h6"
          noWrap
          component="a"
          href="/"
          sx={{
            mr: 2,
            fontFamily: 'monospace',
            fontWeight: 700,
            letterSpacing: '.3rem',
            color: 'white',
            textDecoration: 'none',
          }}
        >
          Newsfeed
        </Typography>
      </div>
      <div className="footerComponent__introduction">
        <p>
          With Newsfeed, content comes to you the minute it&apos;s available.
          Follow websites, podcasts, blogs, and newsletters. Enjoy what&apos;s
          important to you, all in one place.
        </p>
      </div>
      <div className="footerComponent__navLinks">
        {pages.map((page) => (
          <Button
            key={page.name}
            sx={{ my: 1, display: 'block' }}
            className="footerComponent__navLinks--link mx-1"
          >
            <Link href={page.link}>{page.name}</Link>
          </Button>
        ))}
      </div>
      <div className="footerComponent__line"></div>
      <div className="footerComponent__copyright mt-4">
        <Link
          className="footerComponent__copyright--text"
          href={GITHUB_LINK}
          target="_blank"
        >
          Github
        </Link>
      </div>
    </div>
  );
}

export default FooterComponent;
