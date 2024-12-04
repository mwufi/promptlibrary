import Link from "next/link";

interface Prompt {
  id: number;
  title: string;
  content: string;
  category: string;
  votes: number;
  created_at: string;
  created_by: string;
}

async function getPrompts(): Promise<{ prompts: Prompt[], error: string | null }> {
  const startTime = performance.now();
  console.log('[Prompts] Starting to fetch prompts from API');
  
  try {
    const res = await fetch('http://localhost:8000/api/prompts', {
      cache: 'no-store'
    });

    console.log(`[Prompts] API Response received - Status: ${res.status} ${res.statusText}`);
    
    if (!res.ok) {
      if (res.status === 404) {
        console.log('[Prompts] No prompts found (404)');
        return { prompts: [], error: null };
      }
      console.error(`[Prompts] API error: ${res.status} ${res.statusText}`);
      throw new Error(`Failed to fetch prompts: ${res.statusText}`);
    }

    const prompts = await res.json();
    const duration = performance.now() - startTime;
    console.log(`[Prompts] Successfully fetched ${prompts.length} prompts in ${duration.toFixed(2)}ms`);
    prompts.forEach((p: Prompt) => {
      console.log(`[Prompts] Prompt loaded - ID: ${p.id}, Title: "${p.title}", Category: "${p.category}"`);
    });
    
    return { prompts, error: null };
  } catch (error) {
    const duration = performance.now() - startTime;
    console.error(`[Prompts] Error fetching prompts after ${duration.toFixed(2)}ms:`, error);
    return { 
      prompts: [], 
      error: error instanceof Error ? error.message : 'Failed to load prompts. Please try again later.' 
    };
  }
}

export default async function PromptsPage() {
  console.log('[PromptsPage] Starting to render page');
  const { prompts, error } = await getPrompts();

  if (error) {
    console.error('[PromptsPage] Error state:', error);
    return (
      <div className="text-center py-12">
        <p className="text-red-500 mb-4">Error: {error}</p>
        <Link
          href="/prompts/new"
          className="inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 bg-primary text-primary-foreground hover:bg-primary/90 h-10 px-4 py-2"
        >
          Create Your First Prompt
        </Link>
      </div>
    );
  }

  console.log(`[PromptsPage] Rendering ${prompts?.length ?? 0} prompts`);
  
  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <h1 className="text-3xl font-bold">Prompts</h1>
        <Link
          href="/prompts/new"
          className="bg-primary text-primary-foreground px-4 py-2 rounded-lg hover:bg-primary/90"
        >
          Create Prompt
        </Link>
      </div>

      {error ? (
        <div className="text-center py-12">
          <p className="text-destructive mb-4">{error}</p>
          <Link
            href="/prompts/new"
            className="inline-block bg-primary text-primary-foreground px-6 py-3 rounded-lg hover:bg-primary/90"
          >
            Create Your First Prompt
          </Link>
        </div>
      ) : prompts?.length === 0 ? (
        <div className="text-center py-12">
          <p className="text-muted-foreground mb-4">No prompts found. Be the first to create one!</p>
          <Link
            href="/prompts/new"
            className="inline-block bg-primary text-primary-foreground px-6 py-3 rounded-lg hover:bg-primary/90"
          >
            Create Your First Prompt
          </Link>
        </div>
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {prompts?.map((prompt) => (
            <Link
              key={prompt.id}
              href={`/prompts/${prompt.id}`}
              className="block p-6 border rounded-lg hover:border-primary transition-colors"
            >
              <h2 className="text-xl font-semibold mb-2">{prompt.title}</h2>
              <p className="text-muted-foreground mb-4 line-clamp-3">
                {prompt.content}
              </p>
              <div className="flex justify-between text-sm text-muted-foreground">
                <span>{prompt.category}</span>
                <span>⬆️ {prompt.votes}</span>
              </div>
            </Link>
          ))}
        </div>
      )}
    </div>
  );
}
