
import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { LoginPage } from './components/pages/Login'
import { createBrowserRouter, RouterProvider } from 'react-router-dom';

const routers = createBrowserRouter([
  {
    path: '/login',
    element: <LoginPage />
  }
])

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <RouterProvider router={routers}></RouterProvider>
  </StrictMode>,
)
