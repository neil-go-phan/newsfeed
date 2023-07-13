import '@fortawesome/fontawesome-svg-core/styles.css';
import '@/styles/globals.scss';
import { SSRProvider } from 'react-bootstrap';
import type { AppProps } from 'next/app';
import CssBaseline from '@mui/material/CssBaseline';
import { CacheProvider, EmotionCache } from '@emotion/react';
import createEmotionCache from '@/helpers/MUISSRhandle';
import type { ReactElement, ReactNode } from 'react';
import type { NextPage } from 'next';
import { ToastContainer } from 'react-toastify';

const clientSideEmotionCache = createEmotionCache();

export interface MyAppProps extends AppProps {
  emotionCache?: EmotionCache;
  Component: NextPageWithLayout;
}

export type NextPageWithLayout<P = {}, IP = P> = NextPage<P, IP> & {
  getLayout?: (page: ReactElement) => ReactNode;
};

export default function App(props: MyAppProps) {
  const { Component, emotionCache = clientSideEmotionCache, pageProps } = props;
  const getLayout = Component.getLayout ?? ((page) => page);
  return (
    <CacheProvider value={emotionCache}>
      <SSRProvider>
          <CssBaseline />
          {getLayout(<Component {...pageProps} />)}
          <ToastContainer />
      </SSRProvider>
    </CacheProvider>
  );
}
