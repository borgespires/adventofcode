
defmodule Main do
    def find_even_division(row)
    def find_even_division([]),
        do: raise "No even division found... Check your input please!!"
    def find_even_division([h|tail]) do
        case find_for(h, tail) do
            nil -> find_even_division(tail)
            solution -> solution
        end
    end
        
    defp find_for(element, list)
    defp find_for(e, [h|_]) when rem(e, h) == 0, do: div(e, h)
    defp find_for(e, [h|_]) when rem(h, e) == 0, do: div(h, e)
    defp find_for(e, [_|tail]), do: find_for(e, tail)
    defp find_for(_, []), do: nil
end

input = File.stream!("input.txt")
    |> Stream.map(&String.trim/1)
    |> Stream.map(&String.split/1)
    |> Stream.map(fn(x) -> Enum.map(x, &String.to_integer/1) end)
    |> Enum.to_list

input
    |> Enum.map(&(Enum.max(&1) - Enum.min(&1)))
    |> Enum.sum
    |> IO.puts

input
    |> Enum.map(&Main.find_even_division/1)
    |> Enum.sum
    |> IO.puts