import Link from "next/link";

export default function Home() {
  return (
    <div className="space-y-12 max-w-5xl mx-auto">
      <section className="text-center py-16">
        <h1 className="text-5xl font-bold mb-6 bg-gradient-to-r from-primary to-blue-600 bg-clip-text text-transparent">
          Welcome to Prompt Library
        </h1>
        <p className="text-xl text-muted-foreground mb-8 max-w-2xl mx-auto">
          Discover, share, and run AI prompts with the community. Find the perfect prompts for your projects.
        </p>
        <div className="flex gap-4 justify-center">
          <Link
            href="/prompts"
            className="bg-primary text-primary-foreground hover:bg-primary/90 px-8 py-3 rounded-lg font-medium transition-colors"
          >
            Browse Prompts
          </Link>
          <Link
            href="/prompts/new"
            className="bg-secondary text-secondary-foreground hover:bg-secondary/90 px-8 py-3 rounded-lg font-medium transition-colors"
          >
            Create Prompt
          </Link>
        </div>
      </section>

      <section className="grid grid-cols-1 md:grid-cols-3 gap-8">
        <div className="p-6 rounded-lg border bg-card hover:bg-accent/50 transition-colors">
          <h3 className="text-xl font-semibold mb-3">Discover</h3>
          <p className="text-muted-foreground">
            Browse through a curated collection of AI prompts created by the community.
          </p>
        </div>
        <div className="p-6 rounded-lg border bg-card hover:bg-accent/50 transition-colors">
          <h3 className="text-xl font-semibold mb-3">Share</h3>
          <p className="text-muted-foreground">
            Create and share your own prompts with others. Help build a better prompt library.
          </p>
        </div>
        <div className="p-6 rounded-lg border bg-card hover:bg-accent/50 transition-colors">
          <h3 className="text-xl font-semibold mb-3">Run</h3>
          <p className="text-muted-foreground">
            Test prompts directly in your browser and see the results in real-time.
          </p>
        </div>
      </section>

      <section className="text-center py-16">
        <h2 className="text-3xl font-bold mb-4">Ready to get started?</h2>
        <p className="text-muted-foreground mb-8">
          Join our community and start exploring AI prompts today.
        </p>
        <Link
          href="/prompts"
          className="bg-primary text-primary-foreground hover:bg-primary/90 px-8 py-3 rounded-lg font-medium inline-block transition-colors"
        >
          Explore Prompts
        </Link>
      </section>
    </div>
  );
}
