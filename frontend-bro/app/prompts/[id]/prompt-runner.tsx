'use client';

import { useState } from 'react';

interface PromptRunnerProps {
  promptId: string;
}

interface Conversation {
  id: number;
  messages: string;
  created_at: string;
  user_id: string;
  prompt_id: number;
}

export default function PromptRunner({ promptId }: PromptRunnerProps) {
  const [input, setInput] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [conversation, setConversation] = useState<Conversation | null>(null);
  const [error, setError] = useState<string | null>(null);

  const runPrompt = async () => {
    setIsLoading(true);
    setError(null);

    try {
      const res = await fetch(`http://localhost:8000/api/prompts/${promptId}/run`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ input }),
      });

      if (!res.ok) {
        throw new Error('Failed to run prompt');
      }

      const data = await res.json();
      setConversation(data);
      setInput('');
    } catch (err) {
      setError(err instanceof Error ? err.message : 'An error occurred');
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="space-y-4">
      <div className="flex gap-2">
        <input
          type="text"
          value={input}
          onChange={(e) => setInput(e.target.value)}
          placeholder="Enter your input..."
          className="flex-1 px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-primary"
          disabled={isLoading}
        />
        <button
          onClick={runPrompt}
          disabled={isLoading || !input.trim()}
          className="px-4 py-2 bg-primary text-primary-foreground rounded-lg hover:bg-primary/90 disabled:opacity-50"
        >
          {isLoading ? 'Running...' : 'Run'}
        </button>
      </div>

      {error && (
        <div className="p-4 border border-destructive text-destructive rounded-lg">
          {error}
        </div>
      )}

      {conversation && (
        <div className="p-4 border rounded-lg space-y-2">
          <h3 className="font-semibold">Output:</h3>
          <pre className="whitespace-pre-wrap">{conversation.messages}</pre>
        </div>
      )}
    </div>
  );
}
