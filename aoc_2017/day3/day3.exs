
defmodule Main do
    use Bitwise

    def spiral_center(square) do
        get_spiral_size(square)
            |> (&({ div(&1, 2), div(&1, 2)})).()
    end

    def where_is(square) do
        spiral_size = get_spiral_size(square)

        br_corner = spiral_size * spiral_size
        bl_corner = br_corner - (spiral_size - 1)
        tl_corner = bl_corner - (spiral_size - 1)
        tr_corner = tl_corner - (spiral_size - 1)
        
        cond do
            square > bl_corner -> {(square - bl_corner), spiral_size - 1}
            square > tl_corner -> {0, (square - tl_corner)}
            square > tr_corner -> {(square - tr_corner), 0}
            true -> {spiral_size - 1, (tr_corner - square)}
        end
    end

    # smallest odd integer strictly larger than 'square'
    defp get_spiral_size(square) do
        sqrt = :math.sqrt(square)
        ceil = round(:math.ceil(sqrt))
        bor(ceil, 1)
    end

    def manhattan_distance({x1, y1}, {x2, y2}), 
        do: abs(x1 - x2) + abs(y1 - y2)
end

input = 361527
# Main.manhattan_distance(Main.spiral_center(input), Main.where_is(input))
#     |> IO.puts


Main.where_is(5)
    |> IO.inspect