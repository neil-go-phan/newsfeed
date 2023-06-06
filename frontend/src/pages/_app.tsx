import '@fortawesome/fontawesome-svg-core/styles.css';
import '@/styles/globals.scss';
import ProgressBar from '@/common/processBar';
import type { AppProps } from 'next/app';

export default function App({ Component, pageProps }: AppProps) {
  return (
    <>
      <ProgressBar />
      <Component {...pageProps} />
    </>
  );
}
