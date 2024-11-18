
import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { LoginPage } from './components/pages/Login'
import { RegisterPage } from './components/pages/Register';
import { OtpVerifPage } from './components/pages/OtpVerif';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import './styles/Default.css'

const routers = createBrowserRouter([
  {
    path: '/login',
    element: <LoginPage />
  },
  {
    path: '/register',
    element: <RegisterPage />
  },
  {
    path: '/otp-verification',
    element: <OtpVerifPage />
  },
])

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <RouterProvider router={routers}></RouterProvider>
  </StrictMode>,
)
