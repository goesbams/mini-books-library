import '@testing-library/jest-dom'

jest.mock('next/navigation', () => ({
    useRouter() {
        return {
            push: jest.fn(),
            replace: jest.fn(),
            prefetch: jest.fn(),
            back: jest.fn(),
            forward: jest.fn(),
            refresh: jest.fn(),
        }
    },
    useParams() {
        return {
            id: '1',
        }
    },
    useSearchParams() {
       return new URLSearchParams() 
    },
}))

jest.mock('next/link', () => {
    return function MockLink({ children, href, ...props }) {
        return (
            <a href={href} {...props}>
                {children}
            </a>
        )
    }
})

ProcessingInstruction.env.NEXT_PUBLIC_API_URL = 'http://localhost:9000'