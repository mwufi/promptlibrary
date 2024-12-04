'use client';

import Link from 'next/link';
import { usePathname } from 'next/navigation';

export function Navbar() {
  const pathname = usePathname();

  return (
    <nav className="sticky top-0 z-50 w-full border-b bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60">
      <div className="container flex h-14 items-center">
        <div className="mr-4 hidden md:flex">
          <Link href="/" className="mr-6 flex items-center space-x-2">
            <span className="hidden font-bold sm:inline-block">
              Prompt Library
            </span>
          </Link>
          <nav className="flex items-center space-x-6 text-sm font-medium">
            <Link
              href="/prompts"
              className={`transition-colors hover:text-foreground/80 ${
                pathname.startsWith('/prompts') ? 'text-foreground' : 'text-foreground/60'
              }`}
            >
              Browse
            </Link>
            <Link
              href="/prompts/new"
              className={`transition-colors hover:text-foreground/80 ${
                pathname === '/prompts/new' ? 'text-foreground' : 'text-foreground/60'
              }`}
            >
              Create
            </Link>
            <Link
              href="/about"
              className={`transition-colors hover:text-foreground/80 ${
                pathname === '/about' ? 'text-foreground' : 'text-foreground/60'
              }`}
            >
              About
            </Link>
          </nav>
        </div>
        <div className="flex flex-1 items-center justify-between space-x-2 md:justify-end">
          <div className="w-full flex-1 md:w-auto md:flex-none">
            {/* Add search functionality here later */}
          </div>
          <nav className="flex items-center">
            {/* Add user menu here later */}
          </nav>
        </div>
      </div>
    </nav>
  );
}
