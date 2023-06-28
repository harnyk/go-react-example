import { useQuery } from 'react-query';

function App() {
    const { data } = useQuery('helloWorld', () =>
        fetch('/api').then((res) => res.text())
    );

    return (
        <>
            <div>ololo</div>
            <pre>{JSON.stringify(data)}</pre>
        </>
    );
}

export default App;
