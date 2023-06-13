import React, {
  PropsWithChildren,
  useCallback,
  useEffect,
  useState,
} from 'react';
import { useResizeDetector } from 'react-resize-detector';
import Head from 'next/head';
import Sidebar, { SidebarOverlay } from './Sidebar/Sidebar';
import Header from '@/layouts/AdminLayout/Header/Header';
import Footer from '@/layouts/AdminLayout/Footer/Footer';
import { Container } from 'react-bootstrap';
import { checkAuth } from '@/helpers/checkAuth';
import { useRouter } from 'next/router';
import { _ROUTES } from '@/helpers/constants';
import { IsLoggedContext } from '@/common/contexts/isLoggedContext';
import ProgressBar from '@/common/processBar';

export default function AdminLayout({ children }: PropsWithChildren) {
  const [isShowSidebar, setIsShowSidebar] = useState(false);
  const [isLogged, setIsLogged] = useState<boolean>(false);
  const router = useRouter();
  const [isShowSidebarMd, setIsShowSidebarMd] = useState(true);

  const toggleIsShowSidebar = () => {
    setIsShowSidebar(!isShowSidebar);
  };

  const toggleIsShowSidebarMd = () => {
    const newValue = !isShowSidebarMd;
    localStorage.setItem('isShowSidebarMd', newValue ? 'true' : 'false');
    setIsShowSidebarMd(newValue);
  };

  const resetIsShowSidebar = () => {
    setIsShowSidebar(false);
  };

  const onResize = useCallback(() => {
    resetIsShowSidebar();
  }, []);

  const { ref } = useResizeDetector({ onResize });

  useEffect(() => {
    if (localStorage.getItem('isShowSidebarMd')) {
      setIsShowSidebarMd(localStorage.getItem('isShowSidebarMd') === 'true');
    }
  }, [setIsShowSidebarMd]);


  useEffect(() => {
    async function checkLogIn() {
      const userChecked: boolean = await checkAuth();
      setIsLogged(userChecked);
      if (!userChecked) {
        router.push(_ROUTES.LOGIN_PAGE);
      }
    }

    checkLogIn();
    // eslint-disable-next-line
  }, [isLogged]);
  if (isLogged) {
    return (
      <IsLoggedContext.Provider value={{ isLogged, setIsLogged }}>
        <Head>
          <title>Admin page</title>
          <meta name="description" content="Generated by create next app" />
          {/* <meta http-Equiv="Content-Security-Policy" content="upgrade-insecure-requests" /> */}
          <link rel="icon" href="/favicon.ico" />
        </Head>
  
        <div ref={ref} className="position-absolute w-100" />
  
        <Sidebar isShow={isShowSidebar} isShowMd={isShowSidebarMd} />
  
        <div className="wrapper d-flex flex-column min-vh-100 bg-light">
          <Header
            toggleSidebar={toggleIsShowSidebar}
            toggleSidebarMd={toggleIsShowSidebarMd}
          />
          <div className="body flex-grow-1 px-3">
            <Container fluid="lg">{children}</Container>
          </div>
          <Footer />
        </div>
  
        <SidebarOverlay
          isShowSidebar={isShowSidebar}
          toggleSidebar={toggleIsShowSidebar}
        />
      </IsLoggedContext.Provider>
    );
  }
  return (
    <IsLoggedContext.Provider value={{ isLogged, setIsLogged }}>
      <Head>
        <title>Admin page</title>
        <meta name="description" content="Generated by create next app" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <div ref={ref} className="position-absolute w-100" />
      <ProgressBar />
    </IsLoggedContext.Provider>
  );
}
