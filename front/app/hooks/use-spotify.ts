import { useCallback, useEffect, useState } from "react";

interface Props {
  input: string;
  // TODO: add debounce
  debounce: number;
}

export const useSpotify = ({ input }: Props) => {
  const [items, setItems] = useState([]);
  const [hasMore, setHasMore] = useState(true);
  const [isLoading, setIsLoading] = useState(false);
  const [offset, setOffset] = useState(0);
  const track = "track"; // TODO: allow to search by artist, album, etc.
  const limit = 10;

  const loadTracks = useCallback(
    async (offset: number) => {
      const { signal } = new AbortController();

      try {
        setIsLoading(true);

        const res = await fetch(
          `https://api.spotify.com/v1/search?q=${input}&type=${track}&offset=${offset}&limit=${limit}&market=ES`,
          { signal }
        );

        if (!res.ok) {
          throw new Error("Network response was not ok");
        }

        const json = await res.json();
        console.log("FETCHED DATA: ", json);

        setHasMore(json.next !== null);

        // @ts-expect-error needs to be defined
        setItems((prevItems) => [...prevItems, ...json.results]);
      } catch (error) {
        // @ts-expect-error needs to be defined
        if (error.name === "AbortError") {
          console.log("Fetch aborted");
        } else {
          console.error("There was an error with the fetch operation:", error);
        }
      } finally {
        setIsLoading(false);
      }
    },
    [input]
  );

  useEffect(() => {
    loadTracks(offset);
  }, [offset, loadTracks]);

  const onLoadMore = () => {
    const next = offset + limit;

    setOffset(next);
    loadTracks(next);
  };

  return {
    items,
    hasMore,
    isLoading,
    onLoadMore,
  };
};
