'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';

export default function NewPromptPage() {
  const router = useRouter();
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setIsSubmitting(true);
    setError(null);

    const formData = new FormData(e.currentTarget);
    const data = {
      title: formData.get('title'),
      content: formData.get('content'),
      category: formData.get('category'),
    };

    try {
      const res = await fetch('http://localhost:8000/api/prompts', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      });

      if (!res.ok) {
        throw new Error('Failed to create prompt');
      }

      const prompt = await res.json();
      router.push(`/prompts/${prompt.id}`);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'An error occurred');
      setIsSubmitting(false);
    }
  };

  return (
    <div className="max-w-2xl mx-auto">
      <h1 className="text-3xl font-bold mb-8">Create New Prompt</h1>

      <form onSubmit={handleSubmit} className="space-y-6">
        <div className="space-y-2">
          <label htmlFor="title" className="block font-medium">
            Title
          </label>
          <input
            type="text"
            id="title"
            name="title"
            required
            className="w-full p-2 border rounded-lg bg-background text-foreground"
            disabled={isSubmitting}
          />
        </div>

        <div className="space-y-2">
          <label htmlFor="category" className="block font-medium">
            Category
          </label>
          <input
            type="text"
            id="category"
            name="category"
            required
            className="w-full p-2 border rounded-lg bg-background text-foreground"
            disabled={isSubmitting}
          />
        </div>

        <div className="space-y-2">
          <label htmlFor="content" className="block font-medium">
            Prompt Content
          </label>
          <textarea
            id="content"
            name="content"
            required
            className="w-full min-h-[200px] p-2 border rounded-lg bg-background text-foreground"
            disabled={isSubmitting}
          />
        </div>

        {error && (
          <div className="p-4 border border-destructive text-destructive rounded-lg">
            {error}
          </div>
        )}

        <button
          type="submit"
          disabled={isSubmitting}
          className="w-full py-2 bg-primary text-primary-foreground rounded-lg hover:bg-primary/90 disabled:opacity-50"
        >
          {isSubmitting ? 'Creating...' : 'Create Prompt'}
        </button>
      </form>
    </div>
  );
}
