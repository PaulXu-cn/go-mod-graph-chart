import * as d3 from 'd3';
import CSS from '../css/styles.css';

export default function tree() {
    const width = 4000;
    const height = 2500;
    const svg = d3.select('#app').append('svg').attr('id', 'svg')
    .attr('width', width).attr('height', height);
    const margin = { top: 5, right: 150, bottom: 5, left: 50 };
    const innerWidth = width - margin.left - margin.right;
    const innerHeight = height - margin.top - margin.bottom;
    const g = svg.append('g')
        .attr('transform', `translate(${margin.left}, ${margin.top})`);
    let root;
    let color;

    const fill = d => {
        if (d.depth === 0)
            return color(d.data.name)
        while (d.depth > 1)
            d = d.parent;
        return color(d.data.name);
    }

    const render = function (data) {
        color = d3.scaleOrdinal()
            .domain(root.descendants().filter(d => d.depth <= 1).map(d => d.data.name))
            .range(d3.schemeCategory10);

        g.selectAll("path")
            .data(root.links())
            .join("path")
            .attr("fill", "none")
            .attr("stroke", "black")
            .attr("stroke-width", 1.5)
            .attr("d", d3.linkHorizontal().x(d => d.y).y(d => d.x));

        // alternatively, we can use the following code to do a single data-join; 
        // then, use .append() to add circles and texts; 
        /*const node = g.append("g")
        .selectAll("g")
        .data(root.descendants())
        .join("g")
        .attr("transform", d => `translate(${d.y},${d.x})`);*/

        g.selectAll('circle').data(root.descendants()).join('circle')
            // optionally, we can use stroke-linejoin to beautify the path connection; 
            //.attr("stroke-linejoin", "round")
            .attr("stroke-width", 3)
            .attr("fill", fill)
            .attr('cx', d => d.y)
            .attr('cy', d => d.x)
            .attr("r", 6);

        g.selectAll('text').data(root.descendants()).join('text')
            .attr("text-anchor", d => d.children ? "end" : "start")
            // note that if d is a child, d.children is undefined which is actually false! 
            .attr('x', d => (d.children ? -6 : 6) + d.y)
            .attr('y', d => d.x + 5)
            .text(d => d.data.name);
    }

    d3.json('/tree.json').then(data => {
        let treeData = data.data;
        root = d3.hierarchy(treeData);
        // alternatively, we can set size of each node; 
        // root = d3.tree().nodeSize([30, width / (root.height + 1)])(root);
        root = d3.tree().size([innerHeight, innerWidth])(root);
        render(root);
    });
}