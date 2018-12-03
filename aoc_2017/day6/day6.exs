
defmodule Main do
    def to_map(input) do
        input
            |> Enum.with_index
            |> Enum.map(fn {v, i} -> {i, v} end)
            |> Enum.into(%{})
    end

    def find_loop(banks), do: find_loop(banks, %{state(banks) => 0}, 1)
    def find_loop(banks, seen_states, it) do
        next_banks = redistribute(banks)
        next_state = state(next_banks)

        case Map.fetch(seen_states, next_state) do
            {_, seen_it} -> { it, it - seen_it }
            _ -> find_loop(next_banks, Map.put(seen_states, next_state, it), it+1)
        end
    end

    defp redistribute(banks) do
        {idx, n_blocks} = max_with_index(banks)
        
        redistribute(
            Map.put(banks, idx, 0),
            idx + 1,
            n_blocks)
    end
    def redistribute(banks, _, 0), do: banks
    def redistribute(banks, idx, n_blocks) do
        bank_idx = rem(idx, map_size(banks))
        updated = Map.put(banks, bank_idx, banks[bank_idx] + 1)
        redistribute(updated, bank_idx+1, n_blocks-1)
    end

    defp max_with_index(list) do
        list |> Enum.max_by(fn {_, v} -> v end)
    end

    defp state(banks) do
        banks
            |> Enum.map(fn {_, v} -> v end)
            |> Enum.join
    end

    def parse(path) do
        raw = File.read!(path)
        Regex.scan(~r{\d+}, raw, capture: :first)
            |> Enum.map(fn [blocks] -> String.to_integer(blocks) end)
    end
end

Main.parse("input.txt")
    |> Main.to_map
    |> Main.find_loop
    |> IO.inspect