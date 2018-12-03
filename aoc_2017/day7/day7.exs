
defmodule Program do
    defstruct id: nil, weight: nil, parent: nil, children: []

    def from_string(s) do
        reg = ~r{([a-z]*)\s\((\d*)\)(?:\s->\s(.*))?}
        case Regex.run(reg, s, capture: :all_but_first) do
            [id, weight] -> 
                %Program{
                    id: id, 
                    weight: String.to_integer(weight)
                }
            [id, weight, children] -> 
                %Program{
                    id: id, 
                    weight: String.to_integer(weight),
                    children: String.split(children, ", ", trim: true)
                }
        end 
    end

    def get_children(node, tree) do
        node.children
            |> Enum.map(&(tree[&1]))
    end

    def weight(node, tree) do
        children = get_children(node, tree)

        case weight_children(children, tree) do
            {:balanced, weights} -> node.weight + Enum.sum(weights)
            {:unbalanced, weights} -> throw_unbalanced(children, weights)
        end
    end

    defp weight_children(children, tree) do
        weights = children
            |> Enum.map(fn x -> weight(x, tree) end)
        
        {is_balanced?(weights), weights}
    end

    defp is_balanced?([]), do: :balanced
    defp is_balanced?(weights) do
        if length(weights |> Enum.dedup) == 1, 
            do: :balanced,
            else: :unbalanced
    end

    defp throw_unbalanced(children, weights) do
        throw {
            :unbalanced,
            children 
                |> Enum.map(&(&1.weight))
                |> Enum.zip(weights)
        }
    end
end

defmodule Main do
    def read(file) do
        File.stream!(file)
            |> Stream.map(&String.trim/1)
            |> Enum.to_list
    end

    def to_tree(program_list) do
        program_list
            |> Map.new(fn program -> {program.id, program} end)
            |> fill_parents
    end

    defp fill_parents(programs) do
        programs
            |> Enum.flat_map(&all_children/1)
            |> update_parent(programs)
    end

    defp all_children({id, program}) do
        program.children
            |> Enum.map(&({id, &1}))
    end

    defp update_parent([], programs), do: programs
    defp update_parent([{parent, child}|tail], programs) do
        updated = Map.update!(programs, child, fn _ -> %{programs[child] | parent: parent} end )
        update_parent(tail, updated)
    end
end


tree = Main.read("input.txt")
    |> Enum.map(&Program.from_string/1)
    |> Main.to_tree

[{id, root}] = tree 
    |> Enum.filter(fn {_, p} -> p.parent == nil end)

IO.puts(id)

try do
    Program.weight(root, tree)
        |> IO.inspect(charlists: :as_lists)
catch
    x -> IO.inspect(x)
end