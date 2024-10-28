
import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { LoginPage } from './components/pages/Login'
import { RegisterPage } from './components/pages/Register';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';

const routers = createBrowserRouter([
  {
    path: '/login',
    element: <LoginPage />
  },
  {
    path: '/register',
    element: <RegisterPage />
  }
])

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <RouterProvider router={routers}></RouterProvider>
  </StrictMode>,
)
