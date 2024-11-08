sigma = ["a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "ñ", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", " "]
finalStates = ['{:0>2}'.format(x) for x in range(len(sigma))]
print(*['q{}_'.format(x) for x in finalStates])
print()
print("q")
print("05flavio")

print("q 0 -> q0 _ R")
print("q 1 -> q1 _ R")
print("q 2 -> q2 _ R")

qStates = [x for x in range(10)]
for qState in qStates:
    print(f"q0 {qState} -> q0{qState} _ R")
    print(f"q0{qState} _ -> q0{qState}_ _ R")
    for sigIdx in range(len(sigma)):
        print(f"q0{qState} {sigma[sigIdx]} -> q0{qState} {sigma[(sigIdx + qState) % len(sigma)]} R")

    print(f"q1 {qState} -> q1{qState} _ R")
    print(f"q1{qState} _ -> q1{qState}_ _ R")
    for sigIdx in range(len(sigma)):
        print(f"q1{qState} {sigma[sigIdx]} -> q1{qState} {sigma[(sigIdx + qState + 10) % len(sigma)]} R")

qPlusStates = [x for x in range(len(sigma) - 20 )]
for qState in qPlusStates:
    print(f"q2 {qState} -> q2{qState} _ R")
    print(f"q2{qState} _ -> q2{qState}_ _ R")
    for sigIdx in range(len(sigma)):
        print(f"q2{qState} {sigma[sigIdx]} -> q2{qState} {sigma[(sigIdx + qState + 20) % len(sigma)]} R")
