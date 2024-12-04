import { notFound } from "next/navigation";
import PromptRunner from "./prompt-runner";
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

async function getPrompt(id: string): Promise<Prompt | null> {
  const startTime = performance.now();
  console.log(`[PromptDetail] Starting to fetch prompt ID: ${id}`);
  
  try {
    const res = await fetch(`http://localhost:8000/api/prompts/${id}`, {
      cache: 'no-store'
    });

    console.log(`[PromptDetail] API Response received - Status: ${res.status} ${res.statusText}`);

    if (!res.ok) {
      console.error(`[PromptDetail] API error: ${res.status} ${res.statusText}`);
      return null;
    }

    const prompt = await res.json();
    const duration = performance.now() - startTime;
    console.log(`[PromptDetail] Successfully fetched prompt in ${duration.toFixed(2)}ms`);
    console.log(`[PromptDetail] Prompt details - Title: "${prompt.title}", Category: "${prompt.category}", Votes: ${prompt.votes}`);
    
    return prompt;
  } catch (error) {
    const duration = performance.now() - startTime;
    console.error(`[PromptDetail] Error fetching prompt after ${duration.toFixed(2)}ms:`, error);
    return null;
  }
}

export default async function PromptPage({ params }: { params: { id: string } }) {
  console.log(`[PromptPage] Starting to render page for prompt ID: ${params.id}`);
  const prompt = await getPrompt(params.id);

  if (!prompt) {
    console.error(`[PromptPage] Prompt not found - ID: ${params.id}`);
    return (
      <div className="text-center py-12">
        <h1 className="text-2xl font-bold mb-4">Prompt Not Found</h1>
        <p className="text-gray-600 mb-4">The prompt you're looking for doesn't exist.</p>
        <Link
          href="/prompts"
          className="inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 bg-primary text-primary-foreground hover:bg-primary/90 h-10 px-4 py-2"
        >
          Back to Prompts :D
        </Link>
      </div>
    );
  }

  console.log(`[PromptPage] Rendering prompt details - ID: ${params.id}`);

  return (
    <div className="space-y-8">
      <div>
        <h1 className="text-3xl font-bold mb-2">{prompt.title}</h1>
        <div className="flex gap-4 text-sm text-muted-foreground">
          <span>Category: {prompt.category}</span>
          <span>Votes: {prompt.votes}</span>
          <span>Created: {prompt.created_at}</span>
          <span>By: {prompt.created_by}</span>
        </div>
      </div>

      <div className="p-4 border rounded-lg bg-muted/50">
        <pre className="whitespace-pre-wrap">{prompt.content}</pre>
      </div>

      <PromptRunner promptId={params.id} />
    </div>
  );
}
