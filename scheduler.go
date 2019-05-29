package wrr

type Point struct {
	Entry  interface{}
	Weight int
}

type Iterator interface {
	Next() *Point
}

type iterator struct {
	points      []*Point
	divisor     int
	cw          int
	idx         int
	maxWeight   int
	totalWeight int
	cache       []*Point // cache the result after first loop
	cacheIdx    int
}

func (iter *iterator) Next() *Point {
	// read cache
	if len(iter.cache) >= iter.totalWeight {
		iter.cacheIdx = (iter.cacheIdx + 1) % iter.totalWeight

		return iter.cache[iter.cacheIdx]
	}

	// loop from start to end and return the one big or equal cw
	for {
		iter.idx = (iter.idx + 1) % len(iter.points)

		// when loop back to begin, move cw to next step
		if iter.idx == 0 {
			cw := iter.cw - iter.divisor
			// loop end and reset cw to max(S)
			if cw <= 0 {
				cw = iter.maxWeight
				if cw <= 0 {
					return nil
				}
			}
			iter.cw = cw
		}

		if iter.points[iter.idx].Weight >= iter.cw {
			// log cache
			iter.cache = append(iter.cache, iter.points[iter.idx])

			return iter.points[iter.idx]
		}
	}
}

// A [1][2][3][4]
// B [1][2][3][4][5]
// C [1][2]
//
// after initialized
// i = 0, cw = max(S)
//
//               cw
// A [1][2][3][4] |    <- i
// B [1][2][3][4][|5]
// C [1][2]       |
//
// loop from A to C, and return the one big or equal cw (B)
//
// next loop: i = 0 and move cw to next step
//
//             cw
// A [1][2][3][|4]     <- i
// B [1][2][3][|4][5]
// C [1][2]    |
//
// loop from A to C, and return the one big or equal cw (B, A, B)
//
// loop until cw <= 0
//
//   cw
// A |[1][2][3][4]     <- i
// B |[1][2][3][4][5]
// C |[1][2]
//
// reset cw to max(S)
//
// NewWrr weighted round-robin scheduling
func NewWrr(points []*Point) Iterator {
	divisor, maxWeight, totalWeight := 0, 0, 0

	// greatest common divisor of all numbers and cal the max weight and total weight
	for _, point := range points {
		totalWeight += point.Weight

		divisor = gcd(point.Weight, divisor)

		if point.Weight > maxWeight {
			maxWeight = point.Weight
		}
	}

	c := make([]*Point, 0, totalWeight)

	return &iterator{
		points:      points,
		divisor:     divisor,
		cw:          0,
		idx:         -1,
		maxWeight:   maxWeight,
		totalWeight: totalWeight,
		cache:       c,
		cacheIdx:    -1,
	}
}

type smoothIterator struct {
	points      []*Point
	cw          []int
	totalWeight int
	cache       []*Point
	cacheIdx    int
}

func (iter *smoothIterator) Next() *Point {
	// read cache
	if len(iter.cache) >= iter.totalWeight {
		iter.cacheIdx = (iter.cacheIdx + 1) % iter.totalWeight

		return iter.cache[iter.cacheIdx]
	}

	// increase current weight of each point by its weight
	for idx, point := range iter.points {
		iter.cw[idx] += point.Weight
	}

	// select the greatest current weight
	i, _ := max(iter.cw)

	// reduce the selected one's current weight by total weight
	iter.cw[i] -= iter.totalWeight

	// cache
	iter.cache = append(iter.cache, iter.points[i])

	return iter.points[i]
}

func NewSmoothWrr(points []*Point) Iterator {
	cw := make([]int, len(points), len(points))

	totalWeight := 0
	for _, p := range points {
		totalWeight += p.Weight
	}

	c := make([]*Point, 0, totalWeight)

	return &smoothIterator{
		points:      points,
		totalWeight: totalWeight,
		cw:          cw,
		cache:       c,
		cacheIdx:    -1,
	}
}
