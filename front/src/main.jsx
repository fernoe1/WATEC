import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { Canvas } from '@react-three/fiber'
import Experience from './Experience.jsx'
import './index.css';

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <Canvas 
        camera={ {
          fov: 45,
          near: 0.1,
          far: 200,
          position: [ 4, -2, 6 ]
        } }
    >
      <Experience />
    </Canvas>
  </StrictMode>,
)
