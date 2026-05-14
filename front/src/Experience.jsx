import { OrbitControls } from '@react-three/drei';

const Experience = () => {

    return <>
        <OrbitControls makeDefault />

        <mesh scale={ 1.5 } >
            <boxGeometry />
            <meshNormalMaterial />
        </mesh>
    </>
};

export default Experience;