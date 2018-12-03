
defmodule Main do

    def compute(program, calc_offset, p \\ 0, jumps \\ 0)
    def compute(program, _, p, jumps) when p >= map_size(program), do: jumps
    def compute(program, calc_offset, p, jumps) do
        offset = program[p]
        updated_program = Map.put(program, p, calc_offset.(offset))
        compute(updated_program, calc_offset, p + offset, (jumps + 1))
    end

    def to_map(input) do
        input
            |> Enum.with_index
            |> Enum.map(fn {v, i} -> {i, v} end)
            |> Enum.into(%{})
    end
end

input = File.stream!("input.txt")
    |> Stream.map(&String.trim/1)
    |> Stream.map(&String.to_integer/1)
    |> Main.to_map

input
    |> Main.compute(fn o -> o+1 end)
    |> IO.inspect

input
    |> Main.compute(fn o ->
        if o >= 3, do: o-1, else: o+1 end)
    |> IO.inspect