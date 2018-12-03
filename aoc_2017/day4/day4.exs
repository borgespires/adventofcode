
defmodule Main do
    defp valid(pass) do
        pass
            |> Enum.uniq
            |> length
            |> (&(&1==length(pass))).()
    end

    def count_valid(pass_list) do
        pass_list
            |> Enum.map(&valid/1)
            |> Enum.count(&(&1 == true))
    end

    def sort_string(string) do
        string
            |> String.downcase
            |> String.graphemes
            |> Enum.sort
            |> Enum.join
    end
end

input = File.stream!("input.txt")
    |> Stream.map(&String.trim/1)
    |> Stream.map(&String.split/1)
    |> Enum.to_list

input
    |> Main.count_valid
    |> IO.inspect

input
    |> Enum.map(fn(x) -> Enum.map(x, &Main.sort_string/1) end)
    |> Main.count_valid
    |> IO.inspect